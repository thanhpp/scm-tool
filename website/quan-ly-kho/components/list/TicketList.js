import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import InboxIcon from '@mui/icons-material/Inbox';
import Pagination from '../Pagination/Pagination';
import { useDispatch, useSelector } from 'react-redux';
import ItemTicket from '../Item/ItemTicket';
import ItemType from '../Item/ItemType'


function TicketList() {
    const [pagination, setPagination] = useState({
        offset: 0,
        numberPerPage: 2,
        pageCount: 0,
    })

    const [ticketList, setTicketList] = useState()

    const isSearchingSelector = useSelector((state) => state.ticket.isSearching)
    const ticketSelector = useSelector((state) => state.ticket.ticket)


    useEffect(() => {
        const fetchTicket = async () => {
            console.log(ticketSelector)
            try {
                const response = await fetch(' https://scm-tool.thanhpp.ninja/import_ticket')
                if (!response) {
                    throw new Error('somethign wrong');
                    return;
                }
                const data = await response.json()

                setPagination(prev => ({
                    ...prev,
                    pageCount: !isSearchingSelector ? Math.ceil(data.data.length / prev.numberPerPage) : Math.ceil(ticketSelector?.length / prev.numberPerPage)
                }))

                const ticketDisplay = !isSearchingSelector ? data.data.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage)) : ticketSelector?.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage))

                setTicketList(ticketDisplay)
                console.log(ticketSelector)
            } catch (err) {
                console.log('fetch wrong')
            }
        }

        fetchTicket()
    }, [pagination.offset, pagination.numberPerPage, pagination.pageCount, isSearchingSelector, ticketSelector])

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
                <div className=' col-span-1 my-auto   truncate'>supplier id</div>
                <div className=' col-span-1 my-auto   truncate'>storage id</div>
                <div className=' col-span-6 my-auto truncate'>status</div>
                <div className=' col-span-2 my-auto truncate'>fee</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {/* {pagination.currentData ? pagination.currentData.map((item) => <ProductItem isSelectedAll={isSelectedAll} key={item.id} id={item.id} sku={item.sku} name={item.name} category={item.category} price={item.price} amount={item.amount} vendor={item.vendor} arrivalDate={item.arrivalDate} />) : <Box sx={{ display: 'flex' }}> */}
            {ticketList && ticketList.map(ticket => <ItemTicket id={ticket.id} key={ticket.id} supplierId={ticket.from_supplier_id} storageId={ticket.to_storage_id} status={ticket.status} fee={ticket.fee} billImage={ticket.bill_image_paths} productImage={ticket.product_image_paths} receiveTime={ticket.receive_time} sendTime={ticket.send_time} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} />
            {/* <Pagination /> */}
        </div>
    )
}

export default TicketList