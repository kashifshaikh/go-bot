import http from 'k6/http';
import { sleep } from 'k6';

const accessToken = 'ACCESS_TOKEN';

const CognitoUserConfirmEvent = {
  version: 1,
  region: 'us-east-1',
  userPoolId: 'us-east-1_rx2Atvh2H',
  userName: 'c6836649-2318-4f0a-88fe-fa71dc71c011',
  callerContext: {
    awsSdkVersion: 'aws-sdk-unknown-unknown',
    clientId: '58hdos8ji9qa6mchgqd3h7rcnn',
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
      email: 'kashifshaikh2.0+az1@gmail.com',
    },
  },
  response: {},
};

const hello = {
  there: 'yes',
  ok: 1,
  done: true,
};

export default function() {
  const headers = {
    // 'Authorization': `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };
  const body1 = String.raw`{"query":"mutation{ createNewUser(cognitoEvent: \"{\\\"there\\\":\\\"yes\\\",\\\"ok\\\":1,\\\"done\\\":true}\") { id, name, email}}"}`;

  const args = JSON.stringify(CognitoUserConfirmEvent).replace(
    /"/g,
    `\\\\\\\"`,
  );
  const body = String.raw`{"query":"mutation{ createNewUser(cognitoEvent:\"${args}\") { id, name, email}}"}`;
  console.log(body);
  const res = http.post('http://127.0.0.1:8080/graphql', body, {
    headers: headers,
  });
  console.log(res.body);
  // if (res.status === 200) {
  //   // const body = JSON.parse(res.body);
  //   console.log(res.body);
  // }
  // sleep(0.3);
}
