import React, { useState, useEffect } from 'react'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useRouter } from 'next/router';
import Loading from '../../../components/UI/Loading'
import axios from 'axios';
import { useSelector } from 'react-redux';


function GeneralSerial() {
    const [importTicketId, setImportTicketId] = useState()

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


    const importTicketIdHandle = (e) => {
        // const finalData = parseInt(e.target.value)
        setImportTicketId(e.target.value)
    }

    const generalTicketSerial = async () => {
        setLoading(true)
        console.log(JSON.stringify({
            import_ticket_id: parseInt(importTicketId)
        }))
        let token = localStorage.getItem('token')
        try {
            const res = await fetch('https://scm-tool.thanhpp.ninja/import_ticket/serials', {
                method: 'POST',
                body: JSON.stringify({
                    import_ticket_id: parseInt(importTicketId)
                }),
                headers: {
                    'Content-type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            })
            if (!res.ok) {
                throw new Error('something wrong');
                return;
            }

            const data = await res.json()
            if (data.error.code == 200) {
                alert(data.error.message)
                router.push('/io-management/ticket')
            }
            setLoading(false)
            console.log(data)
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
                    General serials
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
                            <label htmlFor='importTicketId'>ticket id</label>
                        </div>
                        <input id='importTicketId' onChange={importTicketIdHandle} type='text' className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div onClick={generalTicketSerial} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                            <button className='text-white ' >import ticket id</button>
                        </div>
                    </div>
                </div>
            </div>
            {loading ? <Loading /> : ''}
        </div>
    )
}

export default GeneralSerial