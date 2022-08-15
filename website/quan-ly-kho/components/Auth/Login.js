import React, { useState } from 'react'
import Loading from '../UI/Loading';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { useSelector, useDispatch } from 'react-redux';
import { authSliceActions } from '../../store/authSlice';

function Login() {
    const [nameuser, setNameuser] = useState()
    const [password, setPassword] = useState()
    const [loading, setLoading] = useState(false)

    const router = useRouter()
    const dispatch = useDispatch()

    const usernameHandle = (e) => {
        setNameuser(e.target.value)
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
                    username: nameuser,
                    password: password
                }),
                headers: {
                    "Content-type": "application/json"
                }
            })

            if (!res.ok) {
                throw new Error('can not login')
                return
            }

            const data = await res.json()
            console.log(data)

            const token = data.data.token
            const username = data.data.user.name
            localStorage.setItem('token', token)
            localStorage.setItem('username', username)
            // dispatch(authSliceActions.setAuth({
            //     data: true
            // }))
            if (data.error.code == 200) {
                alert(data.error.message)
                router.push('/')
            }
            setLoading(false)
        } catch (err) {
            alert(err)
            setLoading(false)
        }

    }
    return (
        <div className=' max-w-[410px] w-full m-auto mt-[200px]'>
            <form onSubmit={submitHandle} className='flex flex-col w-full p-[15px] rounded-lg bg-[#f3f3f3] drop-shadow-md'>
                <p className=' text-center text-[#212509] text-[28px] my-[18px]'>Login</p>
                <div className=' justify-center mb-[5px] '>
                    <input type="text" onChange={usernameHandle} className='w-full p-[10px] border border-gray-600 outline-blue-300' placeholder='Username' />
                    {/* {isValidated.email ?? <p>{isValidated.email}</p>} */}
                </div>
                <div className=' justify-center mb-[10px] '>
                    <input type="password" onChange={passwordHandle} className='w-full p-[10px] border border-gray-600 outline-blue-300' placeholder='Password' />
                    {/* {isValidated.password ?? <p>{isValidated.password}</p>} */}
                </div>
                <button className=' bg-[#70c282] h-[30px] rounded text-white font-medium'>
                    Login
                </button>
                {/* {errorMessage ?? <p>email or password is wrong!</p>} */}
                {/* <Link href='/signup'><p className=' ml-auto mt-3 cursor-pointer hover:opacity-80'>Sign up</p></Link> */}
                {/* <Link href='/signup'><p className=' ml-auto mt-3 cursor-pointer hover:opacity-80'>Change password</p></Link> */}
            </form>
            {loading ? <Loading /> : ''}
        </div>
    )
}


export default Login