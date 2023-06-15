type UserSignUp = {
  email: string
  username: string
  full_name: string
  password: string
}

type UserSignIn = {
  email: string
  password: string
}

export interface UserStateInterface {
  userSignUp: UserSignUp;
  userSignIn: UserSignIn;
  token: string;
  id: string;
  authenticated: boolean;
}

function state(): UserStateInterface {
  return {
    userSignUp: {} as UserSignUp,
    userSignIn: {} as UserSignIn,
    token: '',
    id: '',
    authenticated: false,
  };
}

export default state;
