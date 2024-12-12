// Import Web3
import Web3 from 'web3';

// Initialize Web3
const web3 = new Web3(window.ethereum);

// Request account access
async function connectWallet() {
    try {
        await window.ethereum.request({ method: 'eth_requestAccounts' });
        const accounts = await web3.eth.getAccounts();
        console.log('Connected account:', accounts[0]);
        return accounts[0]; // 0x608cfC1575b56a82a352f14d61be100FA9709D75
    } catch (error) {
        console.error('User  denied account access:', error);
    }
}

// Function to send a transaction
async function sendTransaction(toAddress, amount) {
    const fromAddress = await connectWallet();
    const transactionParameters = {
        to: toAddress, // 0xc14b30303a6736f6e7292aae1ca20708d2e1bfd3
        from: fromAddress, //0x608cfC1575b56a82a352f14d61be100FA9709D75
        value: web3.utils.toHex(web3.utils.toWei(amount, 'ether')), // 19600000000000000000000
    };

    try {
        const txHash = await window.ethereum.request({
            method: 'eth_sendTransaction',
            params: [transactionParameters],
        });
        console.log('Transaction sent with hash:', txHash);
    } catch (error) {
        console.error('Transaction failed:', error);
    }
}
