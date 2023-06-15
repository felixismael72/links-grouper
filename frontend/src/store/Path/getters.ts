import { GetterTree } from 'vuex';
import { StateInterface } from '../index';
import { PathStateInterface } from './state';

const getters: GetterTree<PathStateInterface, StateInterface> = {
  someGetter(/* context */) {
    // your code
  },
};

export default getters;
