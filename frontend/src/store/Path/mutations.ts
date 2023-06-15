import { MutationTree } from 'vuex';
import { PathStateInterface } from './state';

const mutation: MutationTree<PathStateInterface> = {
  setPath(state: PathStateInterface, data) {
    state.paths = [...data.paths];
    state.icons = [...data.icons];
  },
};

export default mutation;
