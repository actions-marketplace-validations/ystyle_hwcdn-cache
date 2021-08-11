# Fresh Huawei Cloud CDN Cache

github action for fresh huawei cloud CDN cache.

## env
- `ACCESS_KEY_ID` ak
- `SECRET_ACCESS_KEY` sk
- `TYPE` file  or directory
- `REGION` default: cn-north-1
- `URLS` split with |
- `ENTERPRISE_PROJECT_ID` enterprise project id


## Example usage
```yaml
uses: ystyle/hwcdn-cache@master
env:
  ACCESS_KEY_ID: id...
  SECRET_ACCESS_KEY: key...
  TYPE: directory
  URLS: https://ystyle.top/|https://ystyle.top/2021/
```

