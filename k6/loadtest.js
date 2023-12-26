import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
    vus: 10, // Number of virtual users
    duration: '30s', // Duration of the test
};

export default function () {
    // Make a GET request to the server
    http.get('http://localhost:1378/');

    // Optionally, pause for a short time
    sleep(1);
}
