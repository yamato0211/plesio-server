import { check } from 'k6';
import ws from 'k6/ws';

export let options = {
    stages: [
        { duration: '1m', target: 20 }, // 1分間で20ユーザーまで増加
        { duration: '1m', target: 40 }, // 次の1分間で40ユーザーまで増加
        { duration: '1m', target: 60 }, // 次の1分間で60ユーザーまで増加
        { duration: '1m', target: 80 }, // 次の1分間で80ユーザーまで増加
        { duration: '1m', target: 100 } // 最後の1分間で100ユーザーまで増加
    ],
    thresholds: {
        ws_connecting: ['p(95)<100'], // 95%の接続が100ms未満であること
    },
};

export default function () {
    const url = 'ws://plesio-api:8000/ws';
    const params = { tags: { my_tag: 'hello' } };

    ws.connect(url, params, function (socket) {
        socket.on('open', function open() {
            console.log('connected');
            let interval = 1000 / 30; // 基本的なメッセージ送信頻度は30fps

            // ステージに応じてメッセージ送信頻度を変更
            if (__ITER >= 20 && __ITER < 40) {
                interval = 1000 / 60; // 60fps
            } else if (__ITER >= 40) {
                interval = 1000 / 120; // 120fps
            }

            socket.setInterval(function timeout() {
                socket.send(Date.now().toString());
            }, interval);
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
