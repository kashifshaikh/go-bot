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
  for (var id = 0; id <= 2; id++) {
    let res = http.post(`http://10.0.0.101:808${id}/graphql`, body, {
      headers: headers
    });
  }
}
