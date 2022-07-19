import http from 'k6/http';
import { sleep } from 'k6';

let accessToken = 'ACCESS_TOKEN';

export default function () {
  let body = {
    name: 'Kashif',
  };

  let headers = {
    // 'Authorization': `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };

  let res = http.post(
    'https://gz-api-serverless-win.azurewebsites.net/api/helloWorld',
    JSON.stringify(body),
    {
      headers: headers,
    }
  );

  if (res.status === 200) {
    // const body = JSON.parse(res.body);
    // console.log(res.body);
  }
  // sleep(0.3);
}
