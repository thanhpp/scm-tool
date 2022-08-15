import { useRouter } from 'next/router'
import React, { useEffect, useLayoutEffect, useState } from 'react'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import Loading from '../../../../components/UI/Loading';
import { useSelector } from 'react-redux';


function TypeId() {
    const [id, setId] = useState()
    const [name, setName] = useState()
    const [desc, setDesc] = useState()
    const [loading, setLoading] = useState(false)
    const router = useRouter()

    const { typeId } = router.query

    const [auth, setAuth] = useState()


    useLayoutEffect(() => {
        if (!localStorage.getItem('token')) {
            router.push('/login')
            return;
        }
        const token = localStorage.getItem('token')
        setAuth(token)
    }, [])


    const fetchProduct = async () => {
        setLoading(true)
        let token = localStorage.getItem('token')
        try {
            const res = await fetch(`https://scm-tool.thanhpp.ninja/item-type`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })

            if (!res.ok) {
                throw new Error('can not fetch data')
                return;
            }

            const data = await res.json()

            const item = data.data.filter((type) => type.id == typeId)
            setId(item[0].id)
            setName(item[0].name)
            setDesc(item[0].desc)

            setProductDetail(...item)
            setImage(item[0].images[0])

            console.log(...item)

            setLoading(false)
        } catch (err) {
            alert(err)
            setLoading(false)
        }
    }

    useEffect(() => {
        if (!router.isReady) {
            return;
        }

        fetchProduct()
    }, [router.isReady])


    const idHandle = (e) => {
        setId(e.target.value)
    }

    const nameHandle = (e) => {
        setName(e.target.value)
    }

    const descHandle = (e) => {
        setDesc(e.target.value)
    }

    const updateTypeHandle = async () => {
        setLoading(true)
        let token = localStorage.getItem('token')
        try {
            const res = await fetch(` https://scm-tool.thanhpp.ninja/item-type/${typeId}`, {
                method: 'PUT',
                body: JSON.stringify({
                    name,
                    desc
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
                router.push('/products-management')
            }
            setLoading(false)
            console.log(data)
        } catch (err) {
            console.log(err)
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
                    {/* {name?.name} */}
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
                            <label htmlFor='id'>Id<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='id' type='text' value={id || ''} onChange={idHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='name'>Name<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='name' type='text' value={name || ''} onChange={nameHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='description'>Description<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='description' type='text' value={desc || ''} onChange={descHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div onClick={updateTypeHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                            <button className='text-white ' >Update type</button>
                        </div>
                    </div>
                </div>
            </div>
            {loading ? <Loading /> : ''}
        </div>
    )
}

export default TypeId