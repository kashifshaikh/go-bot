import http from 'k6/http';
import { sleep } from 'k6';

let accessToken = 'ACCESS_TOKEN';

export default function() {
  let headers = {
    // 'Authorization': `Bearer ${accessToken}`,
    // 'Content-Type': 'application/json',
    Accept: 'application/json'
  };
  let res = http.get(`http://127.0.0.1:8080/graphql?query={user(id:"1"){name}}`, {
    headers: headers
  });
  if (res.status === 200) {
    // const body = JSON.parse(res.body);
    // console.log(res.body);
  }
  // sleep(0.3);
}
