import React, { useState } from 'react'
import LoginIcon from '@mui/icons-material/Login';
import Loading from '../UI/Loading';

function Login() {
    const [username, setUsername] = useState()
    const [password, setPassword] = useState()
    const [loading, setLoading] = useState(false)

    const usernameHandle = (e) => {
        setUsername(e.target.value)
    }

    const passwordHandle = (e) => {
        setPassword(e.target.value)
    }

    const submitHandle = async (e) => {
        e.preventDefault()

        setLoading(true)

        try {
            const res = await fetch(' https://scm-tool.thanhpp.ninja/login', {
                method: "POST",
                body: JSON.stringify({
                    username,
                    password
                }),
                headers: {
                    "Content-ytype": "application/json",
                    "Accept-Encoding": "gzip, deflate, br",
                    "Connection": "keep-alive"
                }
            })

            if (!res.ok) {
                throw new Error('can not login')
                return
            }

            const data = await res.json()
            if (data.error.code == 200) {
                router.push('/')
            }
            setLoading(false)
            console.log(data)
        } catch (err) {
            console.log(err)
            setLoading(false)
        }

    }
    return (
        <div className=' max-w-[410px] w-full m-auto mt-[200px]'>
            <form onSubmit={submitHandle} className='flex flex-col w-full p-[15px] rounded-lg bg-[#f3f3f3] drop-shadow-md'>
                <p className=' text-center text-[#212509] text-[28px] my-[18px]'>Login</p>
                <div className=' justify-center mb-[5px] '>
                    <input type="text" onChange={usernameHandle} className='w-full p-[10px] border border-gray-600 outline-blue-300' placeholder='Email' />
                    {/* {isValidated.email ?? <p>{isValidated.email}</p>} */}
                </div>
                <div className=' justify-center mb-[10px] '>
                    <input type="password" onChange={passwordHandle} className='w-full p-[10px] border border-gray-600 outline-blue-300' placeholder='Password' />
                    {/* {isValidated.password ?? <p>{isValidated.password}</p>} */}
                </div>
                <button className=' bg-[#70c282] h-[30px] rounded text-white font-medium'>
                    <LoginIcon />
                    Login
                </button>
                {/* {errorMessage ?? <p>email or password is wrong!</p>} */}
                <p className=' ml-auto mt-3 cursor-pointer hover:opacity-80'>Quên mật khẩu</p>
            </form>
            {loading ? <Loading /> : ''}
        </div>
    )
}


export default Login