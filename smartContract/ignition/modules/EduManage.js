const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const lINKToken = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
const certNFT = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";

const EduManage = buildModule("eduManage", (m) => {
    // Use the addresses as strings
    const tokenAdd = m.getParameter("tokenAdd", lINKToken);
    const certAdd = m.getParameter("certAdd", certNFT);

    // Assuming the EduManage contract constructor accepts two parameters
    // Make sure the arguments are passed in the correct order
    const eduManage = m.contract("EduManage", [tokenAdd, certAdd]); // Fix the argument order
    return { eduManage };
});

module.exports = EduManage;