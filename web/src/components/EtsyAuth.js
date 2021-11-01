import configData from '../config.json';

const crypto = require("crypto");

async function getAwsSecret(secretName) {
  var AWS = require('aws-sdk'),
    region = configData.AWS_REGION

  AWS.config.update({
    accessKeyId: process.env.REACT_APP_AWS_ACCESS_KEY_ID,
    secretAccessKey: process.env.REACT_APP_AWS_SECRET_ACCESS_KEY
  })

  var client = new AWS.SecretsManager({
    region: region
  });

  return new Promise((resolve, reject) => {
    client.getSecretValue({ SecretId: secretName }, function (err, data) {
      if (err) {
        reject(err);
      }
      else {
        if ('SecretString' in data) {
          resolve(data.SecretString);
        } else {
          let buff = new Buffer(data.SecretBinary, 'base64');
          resolve(buff.toString('ascii'));
        }
      }
    });
  });
}

function HandleRedirect() {
  console.log("Redirected")
}

function base64URLEncode(str) {
  return str
    .toString("base64")
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "")
}

function GenerateCode(codeVerifier) {
  const sha256 = (buffer) => crypto.createHash("sha256").update(buffer).digest();
  const codeChallenge = base64URLEncode(sha256(codeVerifier));
  return codeChallenge
}

function EtsyAuth(scopes) {
  var redirectUrl = "https://www.etsy.com/oauth/connect?response_type=code"
  redirectUrl = `${redirectUrl}&scope=${scopes}`

  var handlerUrl = window.location.href.substring(0, window.location.href.indexOf("/", 9))
  redirectUrl = `${redirectUrl}&redirect_uri=${handlerUrl}/EtsyRedirect`

  var codeVerifier = base64URLEncode(crypto.randomBytes(32));
  var codeChallenge = GenerateCode(codeVerifier);
  redirectUrl = `${redirectUrl}&code_challenge_method=S256`
  redirectUrl = `${redirectUrl}&code_challenge=${codeChallenge}`

  var state = Math.random().toString(36).substring(7);
  redirectUrl = `${redirectUrl}&state=${state}`

  getAwsSecret('dwd/etsyKeystring')
    .then(res => {
      var clientId = JSON.parse(res).Etsy_Keystring;
      redirectUrl = `${redirectUrl}&client_id=${clientId}`
    })
    .then(() => {
      window.location.href = redirectUrl
    })
}

export default EtsyAuth;
