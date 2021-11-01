import configData from '../config.json';
import Cookies from 'universal-cookie';

const crypto = require("crypto");
const cookies = new Cookies();

export default class EtsyAuth {
  static async getAwsSecret(secretName) {
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

  static base64URLEncode(str) {
    return str
      .toString("base64")
      .replace(/\+/g, "-")
      .replace(/\//g, "_")
      .replace(/=/g, "")
  }

  static GenerateCode(codeVerifier) {
    const sha256 = (buffer) => crypto.createHash("sha256").update(buffer).digest();
    const codeChallenge = this.base64URLEncode(sha256(codeVerifier));
    return codeChallenge
  }

  static HandleRedirect() {
    var queryString = window.location.search.substring(1);
    var rawParams = queryString.split("&");
    var params = new Map()
    rawParams.forEach(param => {
      var splitIndex = param.indexOf("=");
      var key = param.substring(0, splitIndex);
      var value = param.substring(splitIndex + 1);
      params.set(key, value);
    })
    var initialState = cookies.get('state');
    if (initialState !== params.get("state")) {
      console.log(`State Mismatch: ${initialState} : ${params.get("state")}`)
      return
    }
    var handlerUrl = window.location.href.substring(0, window.location.href.indexOf("/", 9))
    var codeVer = cookies.get('codeVerifier')
    fetch(`${configData.SERVER_URL}/api/v1/etsy/tokens?redirect_uri=${handlerUrl}/EtsyRedirect&code=${params.get("code")}&code_verifier=${codeVer}`, {
      "method": "GET",
      "headers": {
        "Authorization": configData.API_KEY
      }
    })
      .then(
        (result) => {
          console.log(result)
        })
      .catch(
        (error) => {
          console.log(error)
        }
      )
  }

  static InitAuth(scopes) {
    var redirectUrl = "https://www.etsy.com/oauth/connect?response_type=code"
    redirectUrl = `${redirectUrl}&scope=${scopes}`

    var handlerUrl = window.location.href.substring(0, window.location.href.indexOf("/", 9))
    redirectUrl = `${redirectUrl}&redirect_uri=${handlerUrl}/EtsyRedirect`

    var codeVerifier = this.base64URLEncode(crypto.randomBytes(32));
    cookies.set('codeVerifier', codeVerifier, { path: '/' });

    var codeChallenge = this.GenerateCode(codeVerifier);
    redirectUrl = `${redirectUrl}&code_challenge_method=S256`
    redirectUrl = `${redirectUrl}&code_challenge=${codeChallenge}`

    var state = Math.random().toString(36).substring(7);
    redirectUrl = `${redirectUrl}&state=${state}`
    cookies.set('state', state, { path: '/' });

    this.getAwsSecret('dwd/etsyKeystring')
      .then(res => {
        var clientId = JSON.parse(res).Etsy_Keystring;
        redirectUrl = `${redirectUrl}&client_id=${clientId}`
      })
      .then(() => {
        window.location.href = redirectUrl
      })
  }
}
