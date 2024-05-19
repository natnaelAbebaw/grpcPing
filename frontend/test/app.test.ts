import { PingTestClient } from "../src/ping/PingServiceClientPb";
import { Ping, PingRequest } from "../src/ping/ping_pb";

// i can't mock the server, since i dont get enough information about mocking the grpc server

// describe("ping service", () => {
//   let client: PingTestClient;

//   beforeAll(() => {
//     client = new PingTestClient("http://localhost:8000", {}, {});
//   });

//   it("should return pong", async () => {
//     const request = new PingRequest();

//     request.setPing(new Ping().setText("pong"));

//     const result = await client.ping(request);

//     expect(result.getPing()?.getText()).toEqual("pong");
//   });
// });
