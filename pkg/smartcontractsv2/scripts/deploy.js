async function main() {
    const ScmNFT = await ethers.getContractFactory("ScmNFT")

    // Start deployment, returning a promise that resolves to a contract object
    const scmNFT = await ScmNFT.deploy()
    await scmNFT.deployed()
    console.log("Contract deployed to address:", scmNFT.address)
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })