1. Network Protocol

   1. 네트워크에서 통신을 위한 표준화된 규칙(Protocol)

   - TCP (Transmission Control Protocol) → 신뢰성 있는 연결 기반 통신
   - UDP (User Datagram Protocol) → 빠르지만 신뢰성이 낮은 비연결형 통신
   - HTTP (Hypertext Transfer Protocol) → 웹 서비스에서 사용하는 애플리케이션 계층 프로토콜
   - HTTP/2 → HTTP의 개선된 버전, 바이너리 프레이밍 및 다중화 지원
   - WebSocket → 양방향 실시간 통신을 지원하는 프로토콜 (기본적으로 HTTP 기반)

2. Network Library & Specific technology

   1. 네트워크 프로토콜을 활용한 데이터 교환 방법(기술 포함)

   - Socket (net 패키지 사용) → TCP, UDP 같은 프로토콜을 다루는 저수준 API (네트워크 프로토콜 자체는 아님)
   - JSON-RPC → RPC(Remote Procedure Call) 방식 중 하나이며, 특정 네트워크 프로토콜에 의존하지 않음 (기본적으로 TCP 위에서 작동)
   - gRPC (Google Remote Procedure Call) → HTTP/2 기반의 고성능 RPC 프레임워크
