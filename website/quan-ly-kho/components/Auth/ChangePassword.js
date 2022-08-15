import React, { useState } from 'react'
import Loading from '../UI/Loading';
import Link from 'next/link';
import { useRouter } from 'next/router'

function ChangePassword() {
    const [name, setName] = useState()
    const [username, setUsername] = useState()
    const [password, setPassword] = useState()
    const [loading, setLoading] = useState(false)

    const router = useRouter()

    const newPasswordHandle = (e) => {
        setName(e.target.value)
    }

    const submitHandle = async (e) => {
        e.preventDefault()

        setLoading(true)

        try {
            const res = await fetch(' https://scm-tool.thanhpp.ninja/signup', {
                method: "POST",
                body: JSON.stringify({
                    name,
                    username,
                    password
                }),
                headers: {
                    "Content-type": "application/json",
                }
            })

            if (!res.ok) {
                throw new Error('can not signup')
                return
            }

            const data = await res.json()
            console.log(data)
            if (data.error.code == 200) {
                alert(data.error.message)
                router.push('/login')
            }
            setLoading(false)
            console.log(data)
        } catch (err) {
            alert(err)
            setLoading(false)
        }

    }
    return (
        <div className=' max-w-[410px] w-full m-auto mt-[200px]'>
            <form onSubmit={submitHandle} className='flex flex-col w-full p-[15px] rounded-lg bg-[#f3f3f3] drop-shadow-md'>
                <p className=' text-center text-[#212509] text-[28px] my-[18px]'>Sign Up</p>
                <div className=' justify-center mb-[5px] '>
                    <input type="text" onChange={newPasswordHandle} className='w-full p-[10px] border border-gray-600 outline-blue-300' placeholder='Name' />
                    {/* {isValidated.email ?? <p>{isValidated.email}</p>} */}
                </div>
                <button className=' bg-[#70c282] h-[30px] rounded text-white font-medium'>
                    Change password
                </button>
                {/* {errorMessage ?? <p>email or password is wrong!</p>} */}
                {/* <Link href='/login'><p className=' ml-auto mt-3 cursor-pointer hover:opacity-80'>login</p></Link> */}
            </form>
            {loading ? <Loading /> : ''}
        </div>
    )
}


export default ChangePassword