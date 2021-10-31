import configData from '../config.json';
import { Component } from 'react';

async function getAwsSecret(secretName) {
  var AWS = require('aws-sdk'),
    region = configData.AWS_REGION,
    secretName = secretName

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

export class EtsyAuth extends Component {
  constructor(props) {
    super();
    this.state = { ...props };
  }
  componentWillMount() {
    getAwsSecret('dwd/etsyKeystring')
      .then(res => {
        console.log(res)
        var clientId = JSON.parse(res).Etsy_Keystring;
        window.location.href = "https://www.etsy.com";
      })
  }

  render() {
    return (<section>Redirecting...</section>);
  }
}

export default EtsyAuth;
