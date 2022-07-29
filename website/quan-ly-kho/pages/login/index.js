import React from 'react'
import Login from '../../components/Auth/Login'
import LoginIcon from '@mui/icons-material/Login';


function LoginPage() {
    return (
        <div>
            <Login />
            <div className='max-w-[250px] w-full mx-auto mt-[30px] '>
                <div className=' text-center  p-[20px] mb-3 cursor-pointer hover:opacity-90'>Dành cho khách hàng</div>
                <div className=' text-center  p-[20px] shadow-shadowCustom font-semibold cursor-pointer hover:opacity-90'>Quét mã sản phẩm</div>
            </div>
        </div>
    )
}

export default LoginPage