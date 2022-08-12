/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RarimocoreChangeKeyECDSA {
  newKey?: string;
  signature?: string;
  creator?: string;
}

export interface RarimocoreChangeKeyEdDSA {
  newKey?: string;
  signature?: string;
  creator?: string;
}

export interface RarimocoreConfirmation {
  height?: string;
  root?: string;
  hashes?: string[];
  sigECDSA?: string;
  sigEdDSA?: string;
  creator?: string;
}

export interface RarimocoreDeposit {
  tx?: string;
  fromChain?: string;
  toChain?: string;
  receiver?: string;
  tokenAddress?: string;
  tokenId?: string;
  creator?: string;
}

export type RarimocoreMsgCreateChangeKeyECDSAResponse = object;

export type RarimocoreMsgCreateChangeKeyEdDSAResponse = object;

export type RarimocoreMsgCreateConfirmationResponse = object;

export type RarimocoreMsgCreateDepositResponse = object;

export type RarimocoreMsgDeleteChangeKeyECDSAResponse = object;

export type RarimocoreMsgDeleteChangeKeyEdDSAResponse = object;

export type RarimocoreMsgDeleteConfirmationResponse = object;

export type RarimocoreMsgDeleteDepositResponse = object;

export type RarimocoreMsgUpdateChangeKeyECDSAResponse = object;

export type RarimocoreMsgUpdateChangeKeyEdDSAResponse = object;

export type RarimocoreMsgUpdateConfirmationResponse = object;

export type RarimocoreMsgUpdateDepositResponse = object;

/**
 * Params defines the parameters for the module.
 */
export interface RarimocoreParams {
  keyECDSA?: string;
  keyEdDSA?: string;
}

export interface RarimocoreQueryAllChangeKeyECDSAResponse {
  changeKeyECDSA?: RarimocoreChangeKeyECDSA[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RarimocoreQueryAllChangeKeyEdDSAResponse {
  changeKeyEdDSA?: RarimocoreChangeKeyEdDSA[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RarimocoreQueryAllConfirmationResponse {
  confirmation?: RarimocoreConfirmation[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RarimocoreQueryAllDepositResponse {
  deposit?: RarimocoreDeposit[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RarimocoreQueryGetChangeKeyECDSAResponse {
  changeKeyECDSA?: RarimocoreChangeKeyECDSA;
}

export interface RarimocoreQueryGetChangeKeyEdDSAResponse {
  changeKeyEdDSA?: RarimocoreChangeKeyEdDSA;
}

export interface RarimocoreQueryGetConfirmationResponse {
  confirmation?: RarimocoreConfirmation;
}

export interface RarimocoreQueryGetDepositResponse {
  deposit?: RarimocoreDeposit;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RarimocoreQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: RarimocoreParams;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title rarimocore/change_key_ecdsa.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryChangeKeyEcdsaAll
   * @summary Queries a list of ChangeKeyECDSA items.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/change_key_ecdsa
   */
  queryChangeKeyEcdsaAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RarimocoreQueryAllChangeKeyECDSAResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/change_key_ecdsa`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChangeKeyEcdsa
   * @summary Queries a ChangeKeyECDSA by index.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/change_key_ecdsa/{newKey}
   */
  queryChangeKeyEcdsa = (newKey: string, params: RequestParams = {}) =>
    this.request<RarimocoreQueryGetChangeKeyECDSAResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/change_key_ecdsa/${newKey}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChangeKeyEdDsaAll
   * @summary Queries a list of ChangeKeyEdDSA items.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/change_key_ed_dsa
   */
  queryChangeKeyEdDsaAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RarimocoreQueryAllChangeKeyEdDSAResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/change_key_ed_dsa`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChangeKeyEdDsa
   * @summary Queries a ChangeKeyEdDSA by index.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/change_key_ed_dsa/{newKey}
   */
  queryChangeKeyEdDsa = (newKey: string, params: RequestParams = {}) =>
    this.request<RarimocoreQueryGetChangeKeyEdDSAResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/change_key_ed_dsa/${newKey}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryConfirmationAll
   * @summary Queries a list of Confirmation items.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/confirmation
   */
  queryConfirmationAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RarimocoreQueryAllConfirmationResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/confirmation`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryConfirmation
   * @summary Queries a Confirmation by index.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/confirmation/{height}
   */
  queryConfirmation = (height: string, params: RequestParams = {}) =>
    this.request<RarimocoreQueryGetConfirmationResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/confirmation/${height}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDepositAll
   * @summary Queries a list of Deposit items.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/deposit
   */
  queryDepositAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RarimocoreQueryAllDepositResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/deposit`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDeposit
   * @summary Queries a Deposit by index.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/deposit/{tx}
   */
  queryDeposit = (tx: string, params: RequestParams = {}) =>
    this.request<RarimocoreQueryGetDepositResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/deposit/${tx}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/rarify-protocol/rarimo-core/rarimocore/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<RarimocoreQueryParamsResponse, RpcStatus>({
      path: `/rarify-protocol/rarimo-core/rarimocore/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
