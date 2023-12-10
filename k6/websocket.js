import ws from 'k6/ws';
import { check } from 'k6';

export const options = {
    vus: 50, // 仮想ユーザーの数
    duration: '5m' // テストの継続時間
};

export default function () {
    const url = 'ws://plesio-api:8000/ws';
    const params = { tags: { my_tag: 'hello' } };

    const res = ws.connect(url, params, function (socket) {
        socket.on('open', function open() {
            console.log('connected');
            socket.setInterval(function timeout() {
                socket.send(Date.now().toString());
            }, 1000 / 30); // 30fpsでメッセージを送信
        });

        socket.on('message', function message(data) {
            console.log('Message received: ', data);
        });

        socket.on('close', function close() {
            console.log('disconnected');
        });
    });

    check(res, { 'status is 101': (r) => r && r.status === 101 });
}
