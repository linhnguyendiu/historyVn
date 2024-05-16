const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const initialSupply = 10000000;

const LINKToken = buildModule("LINKToken", (m) => {

    const totalToken = m.getParameter("totalToken", initialSupply);
    const lINKToken = m.contract("LINKToken", [totalToken]);
    return {lINKToken};
}
)

module.exports = LINKToken;