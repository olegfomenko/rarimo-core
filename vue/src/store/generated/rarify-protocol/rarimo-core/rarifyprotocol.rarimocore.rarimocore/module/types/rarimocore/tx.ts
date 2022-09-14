/* eslint-disable */
import { type, typeFromJSON, typeToJSON } from "../tokenmanager/info";
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "rarifyprotocol.rarimocore.rarimocore";

export interface MsgCreateDeposit {
  creator: string;
  index: string;
  tx: string;
  eventId: string;
  fromChain: string;
  tokenType: type;
}

export interface MsgCreateDepositResponse {}

export interface MsgCreateConfirmation {
  creator: string;
  /** hex-encoded */
  root: string;
  indexes: string[];
  /** hex-encoded */
  signatureECDSA: string;
}

export interface MsgCreateConfirmationResponse {}

export interface MsgCreateChangeKeyECDSA {
  creator: string;
  /** hex-encoded */
  newKey: string;
  /** hex-encoded */
  signature: string;
}

export interface MsgCreateChangeKeyECDSAResponse {}

const baseMsgCreateDeposit: object = {
  creator: "",
  index: "",
  tx: "",
  eventId: "",
  fromChain: "",
  tokenType: 0,
};

export const MsgCreateDeposit = {
  encode(message: MsgCreateDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.tx !== "") {
      writer.uint32(26).string(message.tx);
    }
    if (message.eventId !== "") {
      writer.uint32(34).string(message.eventId);
    }
    if (message.fromChain !== "") {
      writer.uint32(42).string(message.fromChain);
    }
    if (message.tokenType !== 0) {
      writer.uint32(48).int32(message.tokenType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDeposit } as MsgCreateDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.tx = reader.string();
          break;
        case 4:
          message.eventId = reader.string();
          break;
        case 5:
          message.fromChain = reader.string();
          break;
        case 6:
          message.tokenType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDeposit {
    const message = { ...baseMsgCreateDeposit } as MsgCreateDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = String(object.tx);
    } else {
      message.tx = "";
    }
    if (object.eventId !== undefined && object.eventId !== null) {
      message.eventId = String(object.eventId);
    } else {
      message.eventId = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = String(object.fromChain);
    } else {
      message.fromChain = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = typeFromJSON(object.tokenType);
    } else {
      message.tokenType = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.tx !== undefined && (obj.tx = message.tx);
    message.eventId !== undefined && (obj.eventId = message.eventId);
    message.fromChain !== undefined && (obj.fromChain = message.fromChain);
    message.tokenType !== undefined &&
      (obj.tokenType = typeToJSON(message.tokenType));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDeposit>): MsgCreateDeposit {
    const message = { ...baseMsgCreateDeposit } as MsgCreateDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.tx !== undefined && object.tx !== null) {
      message.tx = object.tx;
    } else {
      message.tx = "";
    }
    if (object.eventId !== undefined && object.eventId !== null) {
      message.eventId = object.eventId;
    } else {
      message.eventId = "";
    }
    if (object.fromChain !== undefined && object.fromChain !== null) {
      message.fromChain = object.fromChain;
    } else {
      message.fromChain = "";
    }
    if (object.tokenType !== undefined && object.tokenType !== null) {
      message.tokenType = object.tokenType;
    } else {
      message.tokenType = 0;
    }
    return message;
  },
};

const baseMsgCreateDepositResponse: object = {};

export const MsgCreateDepositResponse = {
  encode(
    _: MsgCreateDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateDepositResponse,
    } as MsgCreateDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateDepositResponse {
    const message = {
      ...baseMsgCreateDepositResponse,
    } as MsgCreateDepositResponse;
    return message;
  },

  toJSON(_: MsgCreateDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateDepositResponse>
  ): MsgCreateDepositResponse {
    const message = {
      ...baseMsgCreateDepositResponse,
    } as MsgCreateDepositResponse;
    return message;
  },
};

const baseMsgCreateConfirmation: object = {
  creator: "",
  root: "",
  indexes: "",
  signatureECDSA: "",
};

export const MsgCreateConfirmation = {
  encode(
    message: MsgCreateConfirmation,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.root !== "") {
      writer.uint32(18).string(message.root);
    }
    for (const v of message.indexes) {
      writer.uint32(26).string(v!);
    }
    if (message.signatureECDSA !== "") {
      writer.uint32(34).string(message.signatureECDSA);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateConfirmation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.indexes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.root = reader.string();
          break;
        case 3:
          message.indexes.push(reader.string());
          break;
        case 4:
          message.signatureECDSA = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateConfirmation {
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.indexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = String(object.root);
    } else {
      message.root = "";
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(String(e));
      }
    }
    if (object.signatureECDSA !== undefined && object.signatureECDSA !== null) {
      message.signatureECDSA = String(object.signatureECDSA);
    } else {
      message.signatureECDSA = "";
    }
    return message;
  },

  toJSON(message: MsgCreateConfirmation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.root !== undefined && (obj.root = message.root);
    if (message.indexes) {
      obj.indexes = message.indexes.map((e) => e);
    } else {
      obj.indexes = [];
    }
    message.signatureECDSA !== undefined &&
      (obj.signatureECDSA = message.signatureECDSA);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateConfirmation>
  ): MsgCreateConfirmation {
    const message = { ...baseMsgCreateConfirmation } as MsgCreateConfirmation;
    message.indexes = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.root !== undefined && object.root !== null) {
      message.root = object.root;
    } else {
      message.root = "";
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(e);
      }
    }
    if (object.signatureECDSA !== undefined && object.signatureECDSA !== null) {
      message.signatureECDSA = object.signatureECDSA;
    } else {
      message.signatureECDSA = "";
    }
    return message;
  },
};

const baseMsgCreateConfirmationResponse: object = {};

export const MsgCreateConfirmationResponse = {
  encode(
    _: MsgCreateConfirmationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateConfirmationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateConfirmationResponse,
    } as MsgCreateConfirmationResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateConfirmationResponse {
    const message = {
      ...baseMsgCreateConfirmationResponse,
    } as MsgCreateConfirmationResponse;
    return message;
  },

  toJSON(_: MsgCreateConfirmationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateConfirmationResponse>
  ): MsgCreateConfirmationResponse {
    const message = {
      ...baseMsgCreateConfirmationResponse,
    } as MsgCreateConfirmationResponse;
    return message;
  },
};

const baseMsgCreateChangeKeyECDSA: object = {
  creator: "",
  newKey: "",
  signature: "",
};

export const MsgCreateChangeKeyECDSA = {
  encode(
    message: MsgCreateChangeKeyECDSA,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newKey !== "") {
      writer.uint32(18).string(message.newKey);
    }
    if (message.signature !== "") {
      writer.uint32(26).string(message.signature);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChangeKeyECDSA {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangeKeyECDSA,
    } as MsgCreateChangeKeyECDSA;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newKey = reader.string();
          break;
        case 3:
          message.signature = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateChangeKeyECDSA {
    const message = {
      ...baseMsgCreateChangeKeyECDSA,
    } as MsgCreateChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = String(object.newKey);
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    return message;
  },

  toJSON(message: MsgCreateChangeKeyECDSA): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newKey !== undefined && (obj.newKey = message.newKey);
    message.signature !== undefined && (obj.signature = message.signature);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateChangeKeyECDSA>
  ): MsgCreateChangeKeyECDSA {
    const message = {
      ...baseMsgCreateChangeKeyECDSA,
    } as MsgCreateChangeKeyECDSA;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newKey !== undefined && object.newKey !== null) {
      message.newKey = object.newKey;
    } else {
      message.newKey = "";
    }
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    return message;
  },
};

const baseMsgCreateChangeKeyECDSAResponse: object = {};

export const MsgCreateChangeKeyECDSAResponse = {
  encode(
    _: MsgCreateChangeKeyECDSAResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateChangeKeyECDSAResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateChangeKeyECDSAResponse,
    } as MsgCreateChangeKeyECDSAResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateChangeKeyECDSAResponse {
    const message = {
      ...baseMsgCreateChangeKeyECDSAResponse,
    } as MsgCreateChangeKeyECDSAResponse;
    return message;
  },

  toJSON(_: MsgCreateChangeKeyECDSAResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateChangeKeyECDSAResponse>
  ): MsgCreateChangeKeyECDSAResponse {
    const message = {
      ...baseMsgCreateChangeKeyECDSAResponse,
    } as MsgCreateChangeKeyECDSAResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDeposit(request: MsgCreateDeposit): Promise<MsgCreateDepositResponse>;
  CreateConfirmation(
    request: MsgCreateConfirmation
  ): Promise<MsgCreateConfirmationResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateChangeKeyECDSA(
    request: MsgCreateChangeKeyECDSA
  ): Promise<MsgCreateChangeKeyECDSAResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateDeposit(request: MsgCreateDeposit): Promise<MsgCreateDepositResponse> {
    const data = MsgCreateDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateDeposit",
      data
    );
    return promise.then((data) =>
      MsgCreateDepositResponse.decode(new Reader(data))
    );
  }

  CreateConfirmation(
    request: MsgCreateConfirmation
  ): Promise<MsgCreateConfirmationResponse> {
    const data = MsgCreateConfirmation.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateConfirmation",
      data
    );
    return promise.then((data) =>
      MsgCreateConfirmationResponse.decode(new Reader(data))
    );
  }

  CreateChangeKeyECDSA(
    request: MsgCreateChangeKeyECDSA
  ): Promise<MsgCreateChangeKeyECDSAResponse> {
    const data = MsgCreateChangeKeyECDSA.encode(request).finish();
    const promise = this.rpc.request(
      "rarifyprotocol.rarimocore.rarimocore.Msg",
      "CreateChangeKeyECDSA",
      data
    );
    return promise.then((data) =>
      MsgCreateChangeKeyECDSAResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
