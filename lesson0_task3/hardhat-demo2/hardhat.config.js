require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */

const INFURA_API_ID = "8e1950c1a4f74c34b5741f1e9d518947";
const PRIVATE_KEY = "1cfbd8cb8165a85eae14834992524bdf8cb058fb4d32f814ee3f0fa565e1e775";

module.exports = {
  solidity: "0.8.28",
  // 网络配置
  networks: {
    hardhat: {},
    sepolia: {
      url: "https://sepolia.infura.io/v3/" + `${INFURA_API_ID}`, // 使用Infura时
      // url: `${GLOBAL_API_URL}`, // 使用Alchemy时
      accounts: [`${PRIVATE_KEY}`],
    },
  },
};
