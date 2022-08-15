import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import InboxIcon from '@mui/icons-material/Inbox';
import Pagination from '../Pagination/Pagination';
import { useDispatch, useSelector } from 'react-redux';
import ItemType from '../../components/Item/ItemType'


function ProductList() {
    const [pagination, setPagination] = useState({
        offset: 0,
        numberPerPage: 20,
        pageCount: 0,
    })

    const [itemTypeList, setItemTypeList] = useState()

    const isSearchingSelector = useSelector((state) => state.itemtype.isSearching)
    const itemTypeSelector = useSelector((state) => state.itemtype.itemType)


    useEffect(() => {
        const fetchProduct = async () => {
            let token = localStorage.getItem('token')
            try {
                const response = await fetch(' https://scm-tool.thanhpp.ninja/item-type', {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                })
                if (!response) {
                    throw new Error('somethign wrong');
                    return;
                }
                const data = await response.json()

                setPagination(prev => ({
                    ...prev,
                    pageCount: !isSearchingSelector ? Math.ceil(data.data.length / prev.numberPerPage) : Math.ceil(itemTypeSelector.length / prev.numberPerPage)
                }))

                const warehousesDisplay = !isSearchingSelector ? data.data.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage)) : itemTypeSelector.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage))

                setItemTypeList(warehousesDisplay)
                console.log(data.data)
            } catch (err) {
                alert(err)
            }
        }

        fetchProduct()
    }, [pagination.offset, pagination.numberPerPage, pagination.pageCount, isSearchingSelector, itemTypeSelector])

    const numberPerPageHandle = (value) => {
        setPagination(prev => ({
            ...prev,
            numberPerPage: value
        }))
    }

    const handleClick = (e) => {
        const selected = e.selected
        const offset = selected * pagination.numberPerPage
        setPagination({ ...pagination, offset })
    }
    return (
        //bg-[#323259]
        <div className='bg-white rounded-md'>
            <div className='grid grid-cols-12 gap-3 text-black'>
                <div className='flex items-center col-span-1 my-[16px]'>
                    {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                    <span className='mr-[10px] leading-[24px] ml-[20px]' >id</span>
                </div>
                <div className=' col-span-5 my-auto   truncate'>name</div>
                <div className=' col-span-5 my-auto truncate'>description</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    {/* <span className=' mr-[10px] rounded-sm'>Delete</span> */}
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {/* {pagination.currentData ? pagination.currentData.map((item) => <ProductItem isSelectedAll={isSelectedAll} key={item.id} id={item.id} sku={item.sku} name={item.name} category={item.category} price={item.price} amount={item.amount} vendor={item.vendor} arrivalDate={item.arrivalDate} />) : <Box sx={{ display: 'flex' }}> */}
            {itemTypeList && itemTypeList.map(type => <ItemType id={type.id} key={type.id} name={type.name} desc={type.desc} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} />
            {/* <Pagination /> */}
        </div>
    )
}

export default ProductList