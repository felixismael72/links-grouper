import { ActionTree } from 'vuex';
import { StateInterface } from '../index';
import { UserStateInterface } from './state';
import { auth } from 'src/boot/axios';
import { Notify } from 'quasar';

const actions: ActionTree<UserStateInterface, StateInterface> = {
  signInUser(context) {
    return new Promise((resolve, reject) => {
      auth
        .post('sign-in', context.state.userSignIn)
        .then((response) => {
          context.commit('setAuth', response.data.token);
          Notify.create({
            color: 'green-5',
            textColor: 'white',
            icon: 'check',
            message: 'Sign In executado com sucesso!',
            position: 'top-right',
            timeout: 3000,
          });
          resolve(response);
        })
        .catch((error) => {
          if (error.response) {
            Notify.create({
              color: 'red-5',
              textColor: 'white',
              icon: 'priority_high',
              message: 'Algo deu errado, tente novamente!',
              position: 'top-right',
              timeout: 3000,
            });
          } else {
            Notify.create({
              color: 'red-5',
              textColor: 'white',
              icon: 'priority_high',
              message: 'Oops, algo deu errado! Tente novamente mais tarde.',
              position: 'top-right',
              timeout: 3000,
            });
          }
          reject(error);
        });
    });
  },

  signUpUser(context, userBody) {
    auth
      .post('sign-up', userBody)
      .then((response) => {
        context.commit('setAuth', response.data.token);
        Notify.create({
          color: 'green-5',
          textColor: 'white',
          icon: 'check',
          message: 'Cadastro realizado com sucesso!',
          position: 'top-right',
          timeout: 3000,
        });
      })
      .catch((error) => {
        if (error.response) {
          Notify.create({
            color: 'red-5',
            textColor: 'white',
            icon: 'priority_high',
            message: 'Verifique os dados preenchidos!',
            position: 'top-right',
            timeout: 3000,
          });
        } else {
          Notify.create({
            color: 'red-5',
            textColor: 'white',
            icon: 'priority_high',
            message: 'Oops, algo deu errado! Tente novamente mais tarde.',
            position: 'top-right',
            timeout: 3000,
          });
        }
      })
  }
};

export default actions;
