/* eslint-disable no-undef */
async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);

  const LINKToken = await ethers.getContractFactory("LINKToken");
  const lINKToken = await AnimalsFunding.deploy(1000000000000000000000000);
  await lINKToken.deployed();
  console.log("Contract address:", await  lINKToken.address);

  const CertNFT = await ethers.getContractFactory("CertNFT");
  const certNFT = await CertNFT.deploy();
  await certNFT.deployed();
  console.log("Contract address:", await  certNFT.address);

  const EduManage = await ethers.getContractFactory("EduManage");
  const eduManage = await eduManage.deploy(lINKToken.address);
  await eduManage.deployed();
  console.log("Contract address:", await  eduManage.address);
}

main()
 .then(() => process.exit(0))
 .catch(error => {
   console.error(error);
   process.exit(1);
 });