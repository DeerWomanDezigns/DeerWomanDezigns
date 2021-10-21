import { CognitoUserPool } from 'amazon-cognito-identity-js';


const poolData = {
    UserPoolId: "us-east-2_Fl1juWTiC",
    ClientId: "72bnrc6cdb21mffva2bd5q3ofh",
    RoleArn: 'arn:aws:cognito-idp:us-east-2:394968318459:role/dwd_users',
    AccountId: '394968318459'
}

export default new CognitoUserPool(poolData);