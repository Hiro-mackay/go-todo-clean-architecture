export type Task = {
  id: number;
  title: string;
  createdAt: Date;
  updatedAt: Date;
};

export type CsrfToken = {
  csrf_token: string;
};

export type Credentials = {
  email: string;
  password: string;
};
