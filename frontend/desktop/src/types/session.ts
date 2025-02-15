export type OAuthToken = {
  readonly access_token: string;
  readonly token_type: string;
  readonly refresh_token: string;
  readonly expiry: string;
};

export type UserInfo = {
  readonly k8s_username: string;
  readonly name: string;
  readonly avatar: string;
  readonly nsid: string;
  readonly ns_uid: string;
  readonly userId: string;
};

export type KubeConfig = string;

export type Session = {
  token: string; // jwt token
  // 提供一些简单的信息
  user: UserInfo;
  // 帮忙导出用的
  kubeconfig: KubeConfig;
};
export type JWTPayload = {
  kubeconfig: KubeConfig;
  user: Record<'uid' | 'nsid' | 'k8s_username' | 'ns_uid', string>;
};
export const sessionKey = 'session';
