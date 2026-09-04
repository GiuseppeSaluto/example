import WebSocket from "ws";


// Enables compression, which aisstream.io requires to serve full message bandwidth.
const socket = new WebSocket("wss://stream.aisstream.io/v0/stream", {
    perMessageDeflate: true
})

socket.onopen = function (_) {
    let subscriptionMessage = {
        Apikey: "YOUR API KEY",
        BoundingBoxes: [[[-180, -90], [180, 90]]]
    }
    socket.send(JSON.stringify(subscriptionMessage));
};

socket.onmessage = function (event) {
    const aisMessage = JSON.parse(event.data.toString());
    console.log(aisMessage);
};
