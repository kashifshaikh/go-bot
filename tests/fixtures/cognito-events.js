const SocialUserEvent = {
  version: '1',
  region: 'us-east-1',
  userPoolId: 'us-east-1_rx2Atvh2H',
  userName: 'Google_102769505389790198629',
  callerContext: {
    awsSdkVersion: 'aws-sdk-unknown-unknown',
    clientId: '58hdos8ji9qa6mchgqd3h7rcnn'
  },
  triggerSource: 'PostConfirmation_ConfirmSignUp',
  request: {
    userAttributes: {
      sub: 'ea4d373b-2856-4259-b24c-e818ea8d6256',
      identities:
        '[{"userId":"102769505389790198629","providerName":"Google","providerType":"Google","issuer":null,"primary":true,"dateCreated":1582581501556}]',
      'cognito:user_status': 'EXTERNAL_PROVIDER',
      name: 'Kashif Shaikh',
      email: 'kashifshaikh2.0@gmail.com'
    }
  },
  response: {}
};

const SocialUserEventSignUp = {
  version: '1',
  region: 'us-east-1',
  userPoolId: 'us-east-1_rx2Atvh2H',
  userName: 'Google_102769505389790198629',
  callerContext: {
    awsSdkVersion: 'aws-sdk-unknown-unknown',
    clientId: '58hdos8ji9qa6mchgqd3h7rcnn'
  },
  triggerSource: 'PreSignUp_ExternalProvider',
  request: {
    userAttributes: {
      'cognito:email_alias': '',
      name: 'Kashif Shaikh',
      'cognito:phone_number_alias': '',
      email: 'kashifshaikh2.0@gmail.com'
    },
    validationData: {}
  },
  response: {
    autoConfirmUser: false,
    autoVerifyEmail: false,
    autoVerifyPhone: false
  }
};

const CognitoEvent = {
  version: '1',
  region: 'us-east-1',
  userPoolId: 'us-east-1_rx2Atvh2H',
  userName: 'c6836649-2318-4f0a-88fe-fa71dc71c011',
  callerContext: {
    awsSdkVersion: 'aws-sdk-unknown-unknown',
    clientId: '58hdos8ji9qa6mchgqd3h7rcnn'
  },
  triggerSource: 'PostConfirmation_ConfirmSignUp',
  request: {
    userAttributes: {
      sub: 'c6836649-2318-4f0a-88fe-fa71dc71c011',
      'cognito:email_alias': 'kashifshaikh2.0+az1@gmail.com',
      'cognito:user_status': 'CONFIRMED',
      // eslint-disable-next-line @typescript-eslint/camelcase
      email_verified: 'true',
      name: 'Kashif Shaikh',
      email: 'kashifshaikh2.0+az1@gmail.com'
    }
  },
  response: {}
};

const CognitoSignupEvent = {
  version: '1',
  region: 'us-east-1',
  userPoolId: 'us-east-1_rx2Atvh2H',
  userName: '84587ec3-eb6f-4566-ab42-728a570d6b65',
  callerContext: {
    awsSdkVersion: 'aws-sdk-unknown-unknown',
    clientId: '58hdos8ji9qa6mchgqd3h7rcnn'
  },
  triggerSource: 'PreSignUp_SignUp',
  request: {
    userAttributes: {
      name: 'Kashif Shaikh',
      email: 'kashifshaikh2.0@gmail.com'
    },
    validationData: { invited: 'true' }
  },
  response: {
    autoConfirmUser: false,
    autoVerifyEmail: false,
    autoVerifyPhone: false
  }
};
