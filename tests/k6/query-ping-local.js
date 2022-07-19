import http from 'k6/http';
import { sleep } from 'k6';

let accessToken = 'ACCESS_TOKEN';

export default function() {
  let query = `{ping}`;

  let headers = {
    // 'Authorization': `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
    Accept: 'application/json'
  };

  let body = JSON.stringify({ query: query });
  let res = http.post('http://10.0.0.110:8080/graphql', body, {
    headers: headers
  });

  if (res.status === 200) {
    // const body = JSON.parse(res.body);
    // console.log(res.body);
  }
  // sleep(0.3);
}
