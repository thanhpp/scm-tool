import React, { useState, useRef } from 'react'
import { Filter } from '@mui/icons-material';
import KeyboardArrowDown from '@mui/icons-material/KeyboardArrowDown';

function SearchForm() {

    const [active, setActive] = useState(false)
    const [height, setHeight] = useState('0px')
    const [rotate, setRotate] = useState('transform duration-700 ease')
    const contentSpace = useRef(null)

    const expandAndCollapse = () => {
        setActive(active === false ? true : false)
        setHeight(active ? '0px' : `${contentSpace.current.scrollHeight}px`)
        setRotate(active ? 'transform duration-700 ease' : 'transform duration-700 ease rotate-180')
    }

    return (
        // bg-[#323259]
        <div>
            <div className=' border-white border w-full px-[20px] pt-[20px]'>
                <div className='grid grid-cols-12 gap-3 h-[62px] text-white '>
                    <div className=' col-span-7 '>
                        <input className='w-full border border-white h-[38px] py-[7px] px-[16px] outline-none bg-transparent text-white placeholder-white ' placeholder='Search keywords' type='text' />
                    </div>
                    <div className='col-span-4 flex '>
                        <select id='selectType' defaultValue='Name' className='w-full h-[38px] py-[7px] px-[16px] border border-white outline-none bg-transparent rounded-sm'>
                            <option value='anycategory' className=' text-black'>Name</option>
                            <option value='description' className=' text-black'>Description</option>
                            <option value='location' className=' text-black'>Location</option>
                        </select>
                    </div>
                    {/* <div className='col-span-3'>
                        <select defaultValue='anystock' className=' w-full h-[38px] py-[7px] px-[16px] outline-none bg-transparent border border-white rounded-sm'>
                            <option value="anystock">Any stock status</option>
                            <option value="instock">In stock</option>
                            <option value="lowstock">Low Stock</option>
                            <option value="sold">SOLD</option>
                        </select>
                    </div> */}
                    <div className='col-span-1 h-[38px] text-center hover:opacity-80 rounded-md px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                        <button className='text-white ' >Search</button>
                    </div>
                </div>
                {/* <div className={`flex relative transition-maxHeight duration-700 ease-in-out overflow-hidden`} ref={contentSpace} style={{ maxHeight: `${height}` }}>
                    <div className=' absolute top-[-20px] h-[1px] bg-[#1b1b38] w-[1191px] left-[-20px] '></div>
                    <div className='flex  mr-[40px]'>
                        <div className='mr-[15px]'>Search in:</div>
                        <div className='ml-[10px]'>
                            <ul>
                                <li>
                                    <input type='checkbox' id='name' />
                                    <label htmlFor='name'>Name</label>
                                </li>
                                <li>
                                    <input type='checkbox' id='sku' />
                                    <label htmlFor='sku'>SKU</label>
                                </li>
                                <li className='mb-[8px]'>
                                    <input type='checkbox' id='fulldescription' />
                                    <label htmlFor='fulldescription'>Full Description</label>
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div className='flex mt-[20px] mr-[40px]'>
                        <div className='mr-[15px]'>Availability</div>
                        <div>
                            <select className='w-full h-[38px] py-[7px] px-[16px] outline-none bg-transparent border border-white rounded-sm'>
                                <option value="">Any Availability Status</option>
                                <option value="">Only Enabled</option>
                                <option value="">Only Disabled</option>
                            </select>
                        </div>
                    </div>
                    <div className='flex mt-[20px]'>
                        <div className='mr-[15px]'>Vendor</div>
                        <div>
                            <input className='w-full h-[38px] py-[7px] px-[16px] outline-none bg-transparent border border-white rounded-sm' type='text' />
                        </div>
                    </div>
                </div> */}
            </div>
            {/* <div className='flex justify-center'>
                <div onClick={expandAndCollapse} className=' w-[50px]  text-center cursor-pointer'>
                    <KeyboardArrowDown className={`h-fullr transition-transform ${rotate}`} />
                </div>
            </div> */}
        </div>
    )
}

export default SearchForm