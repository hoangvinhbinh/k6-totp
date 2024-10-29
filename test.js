import totp from 'k6/x/totp'; // Correct import

export default function () {
    const secret = 'JBSWY3DPEHPK3PXP'; // Example TOTP secret (Base32)
    const token = totp.Generate(secret); // Call the method
    console.log(`Generated TOTP: ${token}`); // Log the generated token
}
