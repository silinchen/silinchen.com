import type { Metadata } from 'next'
import { GeistSans } from 'geist/font/sans'
import { GeistMono } from 'geist/font/mono'
import { Analytics } from '@vercel/analytics/react'
import { SpeedInsights } from '@vercel/speed-insights/next'
import Footer from '@/components/footer'
import Navbar from '@/components/nav'
import './globals.css'

const BASE_URL = process.env.BASE_URL as string

export const metadata: Metadata = {
  metadataBase: new URL(BASE_URL),
  title: {
    default: process.env.SITE_NAME as string,
    template: `%s | ${process.env.SITE_NAME} - ${process.env.SITE_DESCRIPTION}`,
  },
  description: process.env.SITE_DESCRIPTION,
  openGraph: {
    title: process.env.SITE_NAME,
    description: process.env.SITE_DESCRIPTION,
    url: BASE_URL,
    siteName: process.env.SITE_NAME,
    locale: 'zh_CN',
    type: 'website',
  },
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      'max-video-preview': -1,
      'max-image-preview': 'large',
      'max-snippet': -1,
    },
  },
}

const cx = (...classes: string[]) => classes.filter(Boolean).join(' ')

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html
      lang="en"
      className={cx(
        'text-black bg-white dark:text-white dark:bg-black',
        GeistSans.variable,
        GeistMono.variable
      )}
    >
      <body className="antialiased max-w-3xl mx-4 mt-8 lg:mx-auto">
        <main className="flex-auto min-w-0 mt-6 flex flex-col px-2 md:px-0">
          <Navbar />
          {children}
          <Footer />
          <Analytics />
          <SpeedInsights />
        </main>
      </body>
    </html>
  )
}
