const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const CertificateNFT = buildModule("CertificateNFT", (m) => {
  const certificateNFT = m.contract("CertificateNFT");

  return { certificateNFT };
});

module.exports = CertificateNFT;