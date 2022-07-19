import http from 'k6/http';
import { sleep } from 'k6';

let accessToken = 'ACCESS_TOKEN';

export default function () {
  let query = `
    query {
      users {
        id
        name
        account {
          id
          name
          domain
        }
      }
    }`;

  let headers = {
    // 'Authorization': `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };

  let res = http.post(
    'https://gz-api-dvamqt5w4q-uc.a.run.app/graphql',
    JSON.stringify({ query: query }),
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
