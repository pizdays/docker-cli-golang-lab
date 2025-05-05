mockery --dir=src/form/domains --name=FormRepository --filename=formRepository_mock.go --output=src/form/mocks/repomocks --outpkg=repomocks

mockery --dir=src/ref/domains --name=RefRepository --filename=refRepository_mock.go --output=src/ref/mocks/repomocks --outpkg=repomocks

mockery --dir=services/database --name=ServicesDatabaseDomain --filename=database_mock.go --output=services/database/mocks --outpkg=servicemocks

mockery --dir=src/operation/domains --name=Repository --filename=Repository_mock.go --output=src/operation/mocks/repomocks --outpkg=repomocks

mockery --dir=src/operation/domains --name=UseCase --filename=UseCase_mock.go --output=src/operation/mocks/usecasemocks --outpkg=usecasemocks