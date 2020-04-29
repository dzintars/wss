# WebSockets

This should be and WSS API gateway for public or private channels.
Ideally it should use some kind of procedure discovery to decide which RPC service to call based on inbound message type.
Outbound messages should be handled by different service if possible by passing WSS connection ID in the context downstream.

https://github.com/dzintars/srp/blob/76194cf379eafa3822ba88e3155fde55a6be8c66/clients/router/static/src/js/redux/actions/messages.js

https://www.youtube.com/watch?v=norUcMSJRtQ

https://exec64.co.uk/blog/websockets_with_redux/

Important article for dynamic json unmarshalling
https://eagain.net/articles/go-json-kind/ (saved locally)
Still need to implement Methods instead of type switches.

## Todo/notes

- Rate limiting for messages, nickname changes, etc
- Organization (chat rooms)
- Teams per organizations

If there is no session cookie, then we need to redirect user to sign in.

If there is an session cookie, then we just open an connection (OR redirect to the private API). Probably we need to track all user sessions because he could use multiple browser tabs or clients.

We should use gorilla secure cookie, because we can easy distribute encryption key across servers via Vault.

We exchange pure actions over wss.

If user are in Order view and there are no other users online in the same order view, then e don't need to create a "chat room" for that order. But as soon as
second user enters the same view, we move booth users in the same connection group???

User roles?

The Hub. What is the Hub? Do i really create a Hub for every single user?

[GitHub Example WSS]()https://github.com/gorilla/websocket/tree/master/examples/chat)

Pool is an object holding all active Hubs.
Every organization with at least one user online have it's own dedicated Hub in that Pool.
Organization Hub is organized in Channels.
Every Channel is a application level View at which every Organization user currently are in.
When last user leaves the current View, that Channel is removed from Hub to clean up the memory and new Channel is created according to the View user navigated to.
If new user joins the View where other user is already on, then new user gets joined to existing Channel and further actions/messages are synced between booth users.
Every user can be part of multiple Organization Channels. For example, he can receive Organization wide events while working in Customer or Order level view/Channel.
Every user can be part of multiple Hubs. For example he can receive System wide events (System itself is an Organization) and Organization wide events.

Main target of this organization is to reduce network traffic and broadcast/receive only what is really required for communications. There is no need to broadcast messages if there is no relative/authorized consumer.

## Articles

https://www.hascode.com/2016/10/writing-a-websocket-chat-in-go/comment-page-1/#comment-7064

[Chat Rooms](https://github.com/gorilla/websocket/issues/46#issuecomment-227906715)

[Go JSON unmarshaling based on an enumerated field value](https://eagain.net/articles/go-json-kind/)

## GitHub
