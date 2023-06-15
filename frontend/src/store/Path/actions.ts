import { ActionTree } from 'vuex';
import { StateInterface } from '../index';
import { PathStateInterface } from './state';

const actions: ActionTree<PathStateInterface, StateInterface> = {
  someAction(/* context */) {
    // your code
  },
};

export default actions;
