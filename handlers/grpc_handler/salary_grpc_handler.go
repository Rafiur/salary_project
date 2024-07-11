package grpchandler

import (
	"context"
	"fmt"
	"salary_project/entity"
	"salary_project/entity/proto"
	"salary_project/service"
)

type SalaryGrpcHandler struct {
	proto.UnimplementedEmployeeToSalaryServer
	SalaryService *service.SalaryService
}

func (s *SalaryGrpcHandler) CreateSalary(ctx context.Context, in *proto.CreateSalaryRequest) (*proto.CreateSalaryResponse, error) {

	fmt.Println(in)
	//fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDD")
	payload := entity.CreateEmployeeSalary{
		Salary_Amount: int(in.SalaryAmount),
		Joining_Date:  in.JoiningDate,
		Project:       in.Project,
		Employee_Id:   in.EmployeeId,
	}
	res, err := s.SalaryService.AddSalaryService(ctx, payload)
	if err != nil {
		fmt.Println(err)
		return &proto.CreateSalaryResponse{}, err
	}

	fmt.Println(res)

	resp := &proto.CreateSalaryResponse{
		Status:  true,
		Message: "Salary Created successfully",
	}

	return resp, nil
}

func (s *SalaryGrpcHandler) GetAllSalary(ctx context.Context, in *proto.GetAllSalaryRequest) (*proto.GetAllSalaryResponse, error) {
	employee_salaries, err := s.SalaryService.GetAllSalaryService(ctx)
	if err != nil {
		return nil, err
	}
	var grpcSalaries []*proto.EmployeeSalary
	for _, salary := range employee_salaries {
		grpcSalaries = append(grpcSalaries, &proto.EmployeeSalary{
			EmployeeId:   int32(salary.Employee_Id),
			SalaryAmount: int32(salary.Salary_Amount),
			Project:      salary.Project,
			JoiningDate:  salary.Joining_Date,
		})
	}
	fmt.Println(grpcSalaries)
	return &proto.GetAllSalaryResponse{Salaries: grpcSalaries}, nil

}

func (s *SalaryGrpcHandler) UpdateSalary(ctx context.Context, in *proto.UpdateSalaryRequest) (*proto.UpdateSalaryResponse, error) {

	//fmt.Println("printing in:",in)

	payload := entity.CreateEmployeeSalary{
		Salary_Amount: int(in.SalaryAmount),
		Joining_Date:  in.JoiningDate,
		Project:       in.Project,
		Employee_Id:   in.EmployeeId,
	}
	//fmt.Println("Printing payload: ",payload)

	res, err := s.SalaryService.UpdateSalaryByIdService(ctx, payload)
	if err != nil {
		fmt.Println(err)
		return &proto.UpdateSalaryResponse{}, err
	}

	fmt.Println(res)

	resp := &proto.UpdateSalaryResponse{
		EmployeeId:   res.Employee_Id,
		SalaryAmount: int32(res.Salary_Amount),
		Project:      res.Project,
		JoiningDate:  res.Joining_Date,
	}

	return resp, nil
}
