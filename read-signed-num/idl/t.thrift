namespace go echo

struct EchoReq {
  1: i64 id,
  2: i32 type,
}

struct EchoResp {
  1: i64 id,
  2: i32 type,
}

service EchoService {
  EchoResp Echo(1: EchoReq req),
}
