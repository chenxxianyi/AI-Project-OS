const qs=(selector,root=document)=>root.querySelector(selector);
const qsa=(selector,root=document)=>[...root.querySelectorAll(selector)];

function toast(message){
  let el=qs('.toast');
  if(!el){
    el=document.createElement('div');
    el.className='toast';
    document.body.appendChild(el);
  }
  el.textContent=message;
  el.classList.add('show');
  clearTimeout(window.__toastTimer);
  window.__toastTimer=setTimeout(()=>el.classList.remove('show'),1800);
}

function openModal(){
  const modal=qs('#createProjectModal');
  if(modal) modal.classList.add('show');
}

function closeModal(){
  const modal=qs('#createProjectModal');
  if(modal) modal.classList.remove('show');
}

qsa('[data-open-modal]').forEach(btn=>btn.addEventListener('click',openModal));
qsa('[data-close-modal]').forEach(btn=>btn.addEventListener('click',closeModal));
qsa('[data-toast]').forEach(btn=>btn.addEventListener('click',()=>toast(btn.dataset.toast||'操作已完成')));
qsa('.modal-backdrop').forEach(backdrop=>backdrop.addEventListener('click',event=>{if(event.target===backdrop)closeModal()}));
qsa('.stack-card').forEach(card=>card.addEventListener('click',()=>{
  const group=card.dataset.group;
  if(group) qsa(`.stack-card[data-group="${group}"]`).forEach(item=>item.classList.remove('selected'));
  card.classList.add('selected');
}));
qsa('[data-next-step]').forEach(btn=>btn.addEventListener('click',()=>toast('已保存技术栈，进入下一步')));
qsa('[data-run-prompt]').forEach(btn=>btn.addEventListener('click',()=>toast('Prompt 已运行，输出已刷新')));

document.addEventListener('keydown',event=>{
  if(event.key==='Escape') closeModal();
  if((event.ctrlKey||event.metaKey)&&event.key==='Enter') toast('Prompt 已运行');
});
