/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

/** Params defines the parameters for the module. */
export interface Params {
  keyECDSA: string;
  keyEdDSA: string;
}

const baseParams: object = { keyECDSA: "", keyEdDSA: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.keyECDSA !== "") {
      writer.uint32(10).string(message.keyECDSA);
    }
    if (message.keyEdDSA !== "") {
      writer.uint32(18).string(message.keyEdDSA);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keyECDSA = reader.string();
          break;
        case 2:
          message.keyEdDSA = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.keyECDSA !== undefined && object.keyECDSA !== null) {
      message.keyECDSA = String(object.keyECDSA);
    } else {
      message.keyECDSA = "";
    }
    if (object.keyEdDSA !== undefined && object.keyEdDSA !== null) {
      message.keyEdDSA = String(object.keyEdDSA);
    } else {
      message.keyEdDSA = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.keyECDSA !== undefined && (obj.keyECDSA = message.keyECDSA);
    message.keyEdDSA !== undefined && (obj.keyEdDSA = message.keyEdDSA);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.keyECDSA !== undefined && object.keyECDSA !== null) {
      message.keyECDSA = object.keyECDSA;
    } else {
      message.keyECDSA = "";
    }
    if (object.keyEdDSA !== undefined && object.keyEdDSA !== null) {
      message.keyEdDSA = object.keyEdDSA;
    } else {
      message.keyEdDSA = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
