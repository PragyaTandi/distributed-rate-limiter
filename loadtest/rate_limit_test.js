import http from 'k6/http';
import { check } from 'k6';
import { Counter } from 'k6/metrics';

// Treat 200 and 429 as expected, so 429 is not counted as a failure
http.setResponseCallback(http.expectedStatuses(200, 429));

const SERVERS = [
  __ENV.SERVER1 || 'http://localhost:8080',
  __ENV.SERVER2 || 'http://localhost:8081',
];
const USERS = ['user1', 'user2', 'user3', 'user4', 'user5'];

const allowed = new Counter('allowed_requests');
const blocked = new Counter('blocked_requests');

export const options = {
  scenarios: {
    steady_load: {
      executor: 'constant-arrival-rate',
      rate: 50,            // 50 requests per second in total
      timeUnit: '1s',
      duration: '30s',
      preAllocatedVUs: 20,
      maxVUs: 50,
    },
  },
  thresholds: {
    http_req_duration: ['p(95)<100'],   // 95% of requests under 100 ms
  },
};

export default function () {
  const server = SERVERS[Math.floor(Math.random() * SERVERS.length)];
  const user = USERS[Math.floor(Math.random() * USERS.length)];

  const res = http.get(server + '/', { headers: { 'X-API-Key': user } });

  check(res, { 'status is 200 or 429': (r) => r.status === 200 || r.status === 429 });

  if (res.status === 200) allowed.add(1);
  if (res.status === 429) blocked.add(1);
}