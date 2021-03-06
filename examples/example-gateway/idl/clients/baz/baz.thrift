namespace java com.uber.zanzibar.clients.baz
include "base.thrift"

enum Fruit {
   APPLE,
   BANANA
}

struct BazRequest {
    1: required bool b1
    2: required string s2
    3: required i32 i3
}

exception AuthErr {
    1: required string message
}

exception OtherAuthErr {
  1: required string message
}

service SimpleService {
    base.BazResponse compare(
        1: required BazRequest arg1
        2: required BazRequest arg2
    ) throws (
        1: AuthErr authErr
        2: OtherAuthErr otherAuthErr
    )

    base.TransStruct trans(
        1: required base.TransStruct arg1
        2: optional base.TransStruct arg2
    ) throws (
        1: AuthErr authErr
        2: OtherAuthErr otherAuthErr
    )

    void call(
        1: required BazRequest arg
    ) throws (
        1: AuthErr authErr
    ) (
        zanzibar.http.reqHeaders = "x-uuid,x-token"
        zanzibar.http.resHeaders = "some-res-header"
    )

    base.BazResponse ping() ()

    void sillyNoop() throws (
        1: AuthErr authErr
        2: base.ServerErr serverErr
    )
}

service SecondService {
    i8 echoI8 (
        1: required i8 arg
    )

    i16 echoI16(
        1: required i16 arg
    )

    i32 echoI32(
        1: required i32 arg
    )

    i64 echoI64(
        1: required i64 arg
    )

    double echoDouble(
        1: required double arg
    )

    bool echoBool (
        1: required bool arg
    )

    binary echoBinary (
        1: required binary arg
    )

    string echoString(
        1: required string arg
    )

    Fruit echoEnum (
        1: required Fruit arg = Fruit.APPLE
    )

    base.UUID echoTypedef(
        1: required base.UUID arg
    )

    set<string> echoStringSet(
        1: required set<string> arg
    )

    // value is unhashable
    set<base.BazResponse> echoStructSet(
        1: required set<base.BazResponse> arg
    )

    list<string> echoStringList (
        1: required list<string> arg
    )

    list<base.BazResponse> echoStructList (
        1: required list<base.BazResponse> arg
    )

    map<string, base.BazResponse> echoStringMap (
        1: required map<string, base.BazResponse> arg
    )

    // key is unhashable
    map<base.BazResponse, string> echoStructMap (
        1: required map<base.BazResponse, string> arg
    )
}
