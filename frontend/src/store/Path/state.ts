export interface PathStateInterface {
  paths: string[];
  icons: string[];
}

function state(): PathStateInterface {
  return {
    paths: [],
    icons: [],
  };
}

export default state;
