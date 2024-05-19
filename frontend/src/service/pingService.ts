import { PingTestClient } from "../ping/PingServiceClientPb";
import { Ping, PingRequest, PingResponse } from "../ping/ping_pb";

export function sendPing(text: string): Promise<PingResponse> {
  const url = "http://localhost:8000";
  const pingService = new PingTestClient(url);
  const ping = new Ping();
  ping.setText(text);

  const pingRequest = new PingRequest();

  pingRequest.setPing(ping);

  return new Promise((resolve, reject) => {
    pingService.ping(pingRequest, {}, function (err, response) {
      if (err) {
        reject(err);
      }
      resolve(response);
    });
  });
}
