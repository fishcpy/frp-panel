'use client'

import { Popover, PopoverTrigger } from '@radix-ui/react-popover'
import { Badge } from '../ui/badge'
import { ClientStatus } from '@/lib/pb/api_master'
import { PopoverContent } from '../ui/popover'
import { useTranslation } from 'react-i18next'
import { motion } from 'framer-motion'
import { formatDistanceToNow } from 'date-fns'
import { zhCN, enUS } from 'date-fns/locale'
import { useStore } from '@nanostores/react'
import { $platformInfo } from '@/store/user'
import { AlertCircle } from 'lucide-react'

export const ClientDetail = ({ clientStatus }: { clientStatus: ClientStatus }) => {
  const { t, i18n } = useTranslation()
  const platformInfo = useStore($platformInfo)

  const locale = i18n.language === 'zh' ? zhCN : enUS
  const connectTime = clientStatus.connectTime
    ? formatDistanceToNow(new Date(parseInt(clientStatus.connectTime.toString())), {
        addSuffix: true,
        locale,
      })
    : '-'

  // 解析版本信息（临时方案：从 githubProxyUrl 字段解析）
  const parseVersionInfo = () => {
    if (!platformInfo?.githubProxyUrl) return null
    const parts = platformInfo.githubProxyUrl.split('|')
    if (parts.length >= 4) {
      return {
        serverVersion: parts[1] || '',
        latestVersion: parts[2] || '',
        enableCheck: parts[3] === 'true',
      }
    }
    return null
  }

  const versionInfo = parseVersionInfo()
  const clientVersion = clientStatus.version?.gitVersion || ''

  // 检查是否需要升级
  const needUpgrade = versionInfo?.enableCheck &&
                      clientVersion &&
                      clientVersion !== 'dev-build' &&
                      versionInfo.latestVersion &&
                      clientVersion !== versionInfo.latestVersion

  return (
    <Popover>
      <PopoverTrigger className="flex items-center gap-1">
        <Badge
          variant="secondary"
          className="text-nowrap rounded-full h-6 hover:bg-secondary/80 transition-colors text-xs"
        >
          {clientStatus.version?.gitVersion || 'Unknown'}
        </Badge>
        {needUpgrade && (
          <AlertCircle className="h-4 w-4 text-yellow-500 animate-pulse" />
        )}
      </PopoverTrigger>
      <PopoverContent className="w-72 p-4 bg-white dark:bg-gray-800 rounded-lg shadow-lg border border-border">
        <motion.div initial={{ opacity: 0, y: -10 }} animate={{ opacity: 1, y: 0 }} transition={{ duration: 0.2 }}>
          <h3 className="text-base font-semibold mb-3 text-center text-foreground">{t('client.detail.title')}</h3>

          {needUpgrade && (
            <div className="mb-3 p-2 bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-800 rounded-md">
              <div className="flex items-center gap-2 text-yellow-800 dark:text-yellow-200">
                <AlertCircle className="h-4 w-4" />
                <span className="text-xs font-medium">
                  {t('client.detail.upgradeAvailable', { version: versionInfo.latestVersion })}
                </span>
              </div>
            </div>
          )}

          <div className="space-y-2">
            <div className="flex justify-between items-center py-1 border-b border-border">
              <span className="text-sm font-medium text-muted-foreground">{t('client.detail.version')}</span>
              <span className="text-sm text-foreground">{clientStatus.version?.gitVersion || '-'}</span>
            </div>
            <div className="flex justify-between items-center py-1 border-b border-border">
              <span className="text-sm font-medium text-muted-foreground">{t('client.detail.buildDate')}</span>
              <span className="text-sm text-foreground">{clientStatus.version?.buildDate || '-'}</span>
            </div>
            <div className="flex justify-between items-center py-1 border-b border-border">
              <span className="text-sm font-medium text-muted-foreground">{t('client.detail.goVersion')}</span>
              <span className="text-sm text-foreground">{clientStatus.version?.goVersion || '-'}</span>
            </div>
            <div className="flex justify-between items-center py-1 border-b border-border">
              <span className="text-sm font-medium text-muted-foreground">{t('client.detail.platform')}</span>
              <span className="text-sm text-foreground">{clientStatus.version?.platform || '-'}</span>
            </div>
            <div className="flex justify-between items-center py-1 border-b border-border">
              <span className="text-sm font-medium text-muted-foreground">{t('client.detail.address')}</span>
              <span className="text-sm text-foreground">{clientStatus.addr || '-'}</span>
            </div>
            <div className="flex justify-between items-center py-1 border-b border-border">
              <span className="text-sm font-medium text-muted-foreground">{t('client.detail.connectTime')}</span>
              <span className="text-sm text-foreground">{connectTime}</span>
            </div>
          </div>
        </motion.div>
      </PopoverContent>
    </Popover>
  )
}
