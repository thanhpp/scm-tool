import React, { useEffect, useState } from 'react'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useRouter } from 'next/router';
import Loading from '../../components/UI/Loading';
import { useSelector } from 'react-redux';


function AddUser() {
    const [name, setName] = useState('')
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [loading, setLoading] = useState(false)
    const router = useRouter()
    const [auth, setAuth] = useState()

    useEffect(() => {
        if (!localStorage.getItem('token')) {
            router.push('/login')
            return;
        }
        const token = localStorage.getItem('token')
        setAuth(token)
    }, [])

    const nameHandle = (e) => {
        setName(e.target.value)
    }

    const usernameHandle = (e) => {
        setUsername(e.target.value)
    }

    const passwordHandle = (e) => {
        setPassword(e.target.value)
    }

    const addUserHandle = async () => {
        setLoading(true)
        const token = localStorage.getItem('token')
        try {
            const res = await fetch('https://scm-tool.thanhpp.ninja/signup', {
                method: 'POST',
                body: JSON.stringify({
                    name,
                    username,
                    password
                }),
                headers: {
                    "Content-type": "application/json",
                    "Authorization": `Bear ${token}`
                }
            })
            if (!res.ok) {
                throw new Error('somethign wrong');
                return;
            }

            const data = await res.json()
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

    return (!auth ? <div></div> :
        <div className='bg-blue-500 w-full h-full min-h-screen   p-[36px]'>
            <div>
                <div onClick={() => router.back()} className=' bg-white rounded-[1000px] w-[30px] h-[30px] mb-[10px] cursor-pointer'>
                    <ArrowBackIcon className='text-black pl-[4px]' />
                </div>
                <div className=' text-[28px]'>
                    Add new user
                </div>
            </div>
            <div className='relative text-[#A16EFF] border-b border-white'>
                <span className='ml-[10px] cursor-pointer'>Info</span>
                <div className=' absolute w-[50px] h-[3px] bg-[#A16EFF]'></div>
            </div>
            <div className='my-[32px]'>
                <div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='name'>Name<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='name' type='text' onChange={nameHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='username'>Usernmae<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='username' type='text' onChange={usernameHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='passwoed'>password<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='passwoed' type='text' onChange={passwordHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div onClick={addUserHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                            <button className='text-white ' >Add new user</button>
                        </div>
                    </div>
                </div>
            </div>
            {loading ? <Loading /> : ''}
        </div>
    )
}

export default AddUser