import { MutationTree } from 'vuex';
import { UserStateInterface } from './state';
import { setLocalToken, removeLocalToken } from 'src/store/utils';

const mutation: MutationTree<UserStateInterface> = {
  fillUser(state: UserStateInterface, data) {
    state.userSignUp = data.user;
  },
  setUserSignIn(state: UserStateInterface, data) {
    state.userSignIn = data
  },
  setAuth(state: UserStateInterface, data) {
    state.token = data;
    const token = state.token;
    state.authenticated = true;

    setLocalToken(token);
    const payload = JSON.parse(atob(token.split('.')[1]));

    state.id = payload.id;
  },
  unsetAuth(state: UserStateInterface) {
    state.userSignIn = { email: '', password: '' };
    state.userSignUp = { email: '', username: '', full_name: '', password: '' };
    (state.authenticated = false),
      (state.token = ''),
      removeLocalToken();
  },
};

export default mutation;
