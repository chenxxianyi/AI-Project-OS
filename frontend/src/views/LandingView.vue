<script setup lang="ts">
import AppButton from '@/components/ui/AppButton.vue'
</script>

<template>
  <div class="landing-body min-h-screen p-[22px_34px_32px] overflow-x-hidden">
    <header class="topbar h-[78px] rounded-[18px] flex items-center justify-between px-6 max-w-[1210px] mx-auto bg-white/72 border border-line shadow-card backdrop-blur-[20px]">
      <router-link to="/" class="flex items-center gap-3 font-black text-lg text-[#0b173d]">
        <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
        AI Project OS
      </router-link>
      <nav class="topnav flex items-center gap-[54px] text-[#425071] font-semibold text-[13px]">
        <router-link to="/dashboard">功能⌄</router-link>
        <router-link to="/dashboard">模板</router-link>
        <router-link to="/dashboard">定价</router-link>
        <router-link to="/dashboard">文档</router-link>
        <router-link to="/dashboard">更新日志</router-link>
      </nav>
      <div class="flex gap-3.5">
        <AppButton variant="ghost"><router-link to="/dashboard">登录</router-link></AppButton>
        <AppButton variant="primary"><router-link to="/dashboard">开始使用</router-link></AppButton>
      </div>
    </header>

    <main>
      <section class="hero grid grid-cols-2 gap-5 max-w-[1120px] mx-auto mt-8 min-h-[492px] items-center relative">
        <div>
          <h1 class="text-[76px] leading-[0.94] tracking-[-0.065em] mb-7">AI Project OS</h1>
          <p class="text-xl leading-[1.8] text-[#62708d] max-w-[620px] mb-8">AI 原生工作台，用于 Prompt 工程、项目上下文、规则管理和多智能体开发。</p>
          <div class="flex gap-5">
            <AppButton variant="primary"><span>✦</span><router-link to="/dashboard">创建项目</router-link></AppButton>
            <AppButton><span>▱</span><router-link to="/dashboard">探索模板</router-link></AppButton>
          </div>
        </div>
        <div class="hero-visual h-[468px] relative grid place-items-center">
          <div class="orb absolute inset-5 bg-[radial-gradient(circle_at_56%_42%,rgba(119,135,255,.32),transparent_27%),radial-gradient(circle_at_76%_46%,rgba(255,196,238,.34),transparent_24%)] blur-[4px]" />
          <div class="cube-stage relative w-[390px] h-[390px] grid place-items-center">
            <div class="cube-platform absolute bottom-[18px] w-[300px] h-[130px] rounded-[48px] [transform:perspective(500px)_rotateX(62deg)_rotateZ(-45deg)] bg-gradient-to-br from-white/68 to-[#d7e3ff]/58 border-2 border-[#c8d6ff]/80 shadow-[0_35px_80px_rgba(59,97,214,.22)]" />
            <div class="cube relative w-[196px] h-[196px] rounded-[42px] bg-gradient-to-br from-white/78 to-[#c3d5ff]/48 border-2 border-[#d3deff]/90 shadow-[inset_0_0_40px_rgba(255,255,255,.9),0_34px_75px_rgba(70,102,210,.27)] rotate-45 grid place-items-center">
              <div class="cube-core w-[42px] h-[42px] rounded-full bg-[radial-gradient(circle,#fff,#8eabff)] shadow-[0_0_38px_rgba(75,105,255,.7)]" />
            </div>
          </div>
          <div class="callout absolute right-0 w-[210px] p-[18px_20px] rounded-2xl bg-white/72 border border-line shadow-card flex gap-3.5 items-center top-[46px]">
            <span class="icon-box">✦</span>
            <span><b class="block text-[13px]">Prompt 工作室</b><span class="text-[11px] text-muted">编写强大的提示词</span></span>
          </div>
          <div class="callout absolute right-[-22px] w-[210px] p-[18px_20px] rounded-2xl bg-white/72 border border-line shadow-card flex gap-3.5 items-center top-[152px]">
            <span class="icon-box">▣</span>
            <span><b class="block text-[13px]">规则构建器</b><span class="text-[11px] text-muted">定义项目规则</span></span>
          </div>
          <div class="callout absolute right-0 w-[210px] p-[18px_20px] rounded-2xl bg-white/72 border border-line shadow-card flex gap-3.5 items-center top-[270px]">
            <span class="icon-box purple">◉</span>
            <span><b class="block text-[13px]">项目大脑</b><span class="text-[11px] text-muted">组织项目知识</span></span>
          </div>
          <div class="callout absolute right-[-18px] w-[210px] p-[18px_20px] rounded-2xl bg-white/72 border border-line shadow-card flex gap-3.5 items-center top-[382px]">
            <span class="icon-box">✹</span>
            <span><b class="block text-[13px]">上下文引擎</b><span class="text-[11px] text-muted">理解项目全貌</span></span>
          </div>
        </div>
      </section>

      <section class="feature-strip max-w-[1118px] mx-auto my-[18px_38px] p-[22px_30px] rounded-[18px] grid grid-cols-5 bg-white/65 border border-line shadow-card">
        <div class="feature-tile text-center min-h-[112px] grid place-items-center border-r border-line">
          <div><span class="icon-box big">✧</span><b class="block my-2.5">Prompt 工作室</b><span class="block text-muted leading-[1.55] text-xs">设计和迭代<br>提示词</span></div>
        </div>
        <div class="feature-tile text-center min-h-[112px] grid place-items-center border-r border-line">
          <div><span class="icon-box big">⬡</span><b class="block my-2.5">规则构建器</b><span class="block text-muted leading-[1.55] text-xs">为 AI 和团队<br>创建规则</span></div>
        </div>
        <div class="feature-tile text-center min-h-[112px] grid place-items-center border-r border-line">
          <div><span class="icon-box big">▱</span><b class="block my-2.5">项目大脑</b><span class="block text-muted leading-[1.55] text-xs">集中管理文档、<br>知识和代码</span></div>
        </div>
        <div class="feature-tile text-center min-h-[112px] grid place-items-center border-r border-line">
          <div><span class="icon-box big">☼</span><b class="block my-2.5">上下文引擎</b><span class="block text-muted leading-[1.55] text-xs">AI 理解你的<br>整个项目</span></div>
        </div>
        <div class="feature-tile text-center min-h-[112px] grid place-items-center">
          <div><span class="icon-box big">♙</span><b class="block my-2.5">智能体工作流</b><span class="block text-muted leading-[1.55] text-xs">构建多智能体<br>应用</span></div>
        </div>
      </section>

      <section class="trusted text-center text-muted max-w-[880px] mx-auto">
        <p>受到构建未来的开发者和团队信赖</p>
        <div class="logos flex justify-between mt-[30px] text-[#2e3854] font-black text-[17px]">
          <span>▲ Vercel</span>
          <span>⬢ Cursor</span>
          <span>◆ Supabase</span>
          <span>▰ Replicate</span>
          <span>◎ OpenAI</span>
        </div>
      </section>
    </main>
  </div>
</template>
