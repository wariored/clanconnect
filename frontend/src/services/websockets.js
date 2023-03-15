// socketService.js
export default class SocketService {
  constructor(url) {
    this.socket = new WebSocket(url, ['Bearer', "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjNmY2JiNzM2MWNiY2RkOWY4Nzc1MGEyIiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNjc4ODc3MDc3fQ.mC8WPk4ZayrAqUueLygQHQMKNXAQUnUfZ1MeALY5GoY"])
    this.socket.onopen = this.onOpen.bind(this)
    this.socket.onmessage = this.onMessage.bind(this)
    this.socket.onclose = this.onClose.bind(this)
  }

  onOpen() {
    console.log('WebSocket connected')
  }

  onMessage(event) {
    console.log('WebSocket message received: ', event.data)
  }

  onClose() {
    console.log('WebSocket closed')
  }

  send(message) {
    this.socket.send(message)
  }

  close() {
    this.socket.close()
  }
}
