import React, { useEffect, useState } from 'react'
import ReactDOM from 'react-dom'
// import classes from './modal.module.css'
import CloseIcon from '@mui/icons-material/Close';

function Modal({ show, onClose, children }) {

  const [isBrowser, setIsBrowser] = useState(false)

  useEffect(() => {
    setIsBrowser(true)
  }, [])

  const handleClose = () => {
    onClose()
  }

  const modalContent = show ? (
    <div className=' overflow-scroll'>
      <div className='backdrop' onClick={handleClose}>Close</div>
      <div className='modal'>
        <div className=' cursor-pointer flex justify-end' onClick={handleClose}><CloseIcon /></div>
        <div >{children}</div>
        {/* <div className='flex '>
                    <div className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px]  border border-white bg-blue-500 cursor-pointer'>
                        <button className='text-black ' >Update product</button>
                    </div>
                </div> */}
      </div>
      <style jsx>{`
                .backdrop {
                    position: fixed;
                    top: 0;
                    left: 0;
                    width: 100%;
                    height: 100vh;
                    z-index: 20;
                    background-color: rgba(0, 0, 0, 0.75);
                  }
                  
                  .modal {
                    position: fixed;
                    top: 8vh;
                    left: 5%;
                    width: 90%;
                    background-color: white;
                    padding: 1rem;
                    border-radius: 14px;
                    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.25);
                    z-index: 30;
                    animation: slide-down 300ms ease-out forwards;
                  }
                  
                  @media (min-width: 768px) {
                    .modal {
                      width: 40rem;
                      left: calc(50% - 20rem);
                    }
                  }
                  
                  @keyframes slide-down {
                    from {
                      opacity: 0;
                      transform: translateY(-3rem);
                    }
                    to {
                      opacity: 1;
                      transform: translateY(0);
                    }
                  }
            `}</style>
    </div>
  ) : null

  if (isBrowser) {
    return ReactDOM.createPortal(
      modalContent,
      document.getElementById('modal-root')
    )
  } else {
    return null
  }
}

export default Modal