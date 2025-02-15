import SigninComponent from '@/components/signin';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';

export default function SigninPage() {
  return <SigninComponent />;
}

export async function getServerSideProps({ req, res, locales }: any) {
  const lang: string = req?.headers?.['accept-language'] || 'zh';
  const local = lang.indexOf('zh') !== -1 ? 'zh' : 'en';
  res.setHeader('Set-Cookie', `NEXT_LOCALE=${local}; Max-Age=2592000; Secure; SameSite=None`);

  const props = {
    ...(await serverSideTranslations(local, undefined, null, locales || []))
  };
  return {
    props
  };
}
